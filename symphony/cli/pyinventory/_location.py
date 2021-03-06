#!/usr/bin/env python3
# pyre-strict

from typing import Any, Dict, List, Optional, Tuple

from gql.gql.client import OperationException

from ._utils import _get_properties_to_add
from .consts import Document, ImageEntity, Location
from .exceptions import (
    LocationCannotBeDeletedWithDependency,
    LocationIsNotUniqueException,
    LocationNotFoundException,
)
from .graphql.add_location_mutation import AddLocationInput, AddLocationMutation
from .graphql.edit_location_mutation import EditLocationInput, EditLocationMutation
from .graphql.location_children_query import LocationChildrenQuery
from .graphql.location_deps_query import LocationDepsQuery
from .graphql.location_details_query import LocationDetailsQuery
from .graphql.location_documents_query import LocationDocumentsQuery
from .graphql.move_location_mutation import MoveLocationMutation
from .graphql.remove_location_mutation import RemoveLocationMutation
from .graphql.search_query import SearchQuery
from .graphql_client import GraphqlClient
from .reporter import FailedOperationException


ADD_LOCATION_MUTATION_NAME = "addLocation"
EDIT_LOCATION_MUTATION_NAME = "editLocation"
MOVE_LOCATION_MUTATION_NAME = "moveLocation"


def get_location(
    client: GraphqlClient, location_hirerchy: List[Tuple[str, str]]
) -> Location:
    """This function returns a location of a specific type with a specific name.
        It can get only the requested location specifiers or the hirerchy leading to it

        Args:
            location_hirerchy (list of tuple(str, str)):
                the first str is location type name
                the second str is location name

        Returns: client.Location object

        Raises: LocationIsNotUniqueException if there is more than one correct
                location to return
                or LocationNotFoundException if no location was found

        Example:
                location = client.getLocation([
                    ('Country', 'England'),
                    ('City', 'Milton Keynes'),
                    ('Site', 'Bletchley Park')
                ])
            or
                location = client.getLocation([('Site', 'Bletchley Park')])
                # This call will fail if there is Bletchley Park in two cities in london
    """

    last_location = None

    for location in location_hirerchy:
        location_type = location[0]
        location_name = location[1]

        if last_location is None:
            locations = SearchQuery.execute(
                client, name=location_name
            ).searchForEntity.edges

            locations = [
                location.node
                for location in locations
                if location.node.entityType == "location"
                and location.node.type == location_type
                and location.node.name == location_name
            ]
            if len(locations) == 0:
                raise LocationNotFoundException(
                    location_name=location_name, location_type=location_type
                )
            if len(locations) != 1:
                raise LocationIsNotUniqueException(
                    location_name=location_name, location_type=location_type
                )
            location_details = LocationDetailsQuery.execute(
                client, id=locations[0].entityId
            ).location
            last_location = Location(
                name=location_details.name,
                id=location_details.id,
                externalId=location_details.externalId,
            )
        else:
            location_id = last_location.id

            result = LocationChildrenQuery.execute(client, id=location_id)
            locations = result.location.children

            locations = [
                location
                for location in locations
                if location.locationType.name == location_type
                and location.name == location_name
            ]
            if len(locations) == 0:
                raise LocationNotFoundException(location_name=location_name)
            if len(locations) != 1:
                raise LocationIsNotUniqueException(
                    location_name=location_name, location_type=location_type
                )
            last_location = Location(
                name=locations[0].name,
                id=locations[0].id,
                externalId=locations[0].externalId,
            )

    if last_location is None:
        raise LocationNotFoundException()
    return last_location


def get_location_children(client: GraphqlClient, location_id: str) -> List[Location]:
    """This function returns all locations that are children of the given location

        Args:
            location_id (str):
                id of the parent location

        Returns: List of client.Location objects

        Example:
                client.addLocation([('Country', 'England'), ('City', 'Milton Keynes')], {})
                client.addLocation([('Country', 'England'), ('City', 'London')], {})
                locations = client.get_location_children([('Country', 'England')])
                # This call will return a list with 2 locations: 'Milton Keynes' and 'London'
    """
    result = LocationChildrenQuery.execute(client, id=location_id)
    locations = result.location.children

    if len(locations) == 0:
        return []

    return [
        Location(name=location.name, id=location.id, externalId=location.externalId)
        for location in locations
    ]


def add_location(
    client: GraphqlClient,
    location_hirerchy: List[Tuple[str, str]],
    properties_dict: Dict[str, Any],
    lat: Optional[float] = None,
    long: Optional[float] = None,
    externalID: Optional[str] = None,
) -> Location:
    """Create a new location of a specific type with a specific name.
        It will also get the requested location specifiers for hirerchy leading to it and will create all
        the hirerchy.
        However the lat,long and propertiesDict would only apply for the last location in the chain.
        If a location with his name in this place already exists the existing location is returned

        Args:
            location_hirerchy (List[Tuple[str, str]]):
                An hirerchy of locations.
                the first str is location type name
                the second str is location name
            properties_dict: dict of property name to property value. the property value should match
                            the property type. Otherwise exception is raised
            lat (float): latitude
            long (float): longitude
            external id (str): ID from external system

        Returns: client.Location object

        Raises: LocationIsNotUniqueException if there is two possible locations
                    inside the chain and it is not clear where to create or
                    what to return
                FailedOperationException for internal inventory error

        Example:
                location = client.add_location(
                    [
                        ('Country', 'England'),
                        ('City', 'Milton Keynes'),
                        ('Site', 'Bletchley Park')
                    ],
                    {
                        'Date Property ': date.today(),
                        'Lat/Lng Property: ': (-1.23,9.232),
                        'E-mail Property ': "user@fb.com",
                        'Number Property ': 11,
                        'String Property ': "aa",
                        'Float Property': 1.23
                    },
                    -11.32,
                    98.32,
                    None)
    """

    last_location = None

    for i, location in enumerate(location_hirerchy):
        location_type = location[0]
        location_name = location[1]

        properties = []
        lat_val = None
        long_val = None
        if i == len(location_hirerchy) - 1:
            property_types = client.locationTypes[location_type].propertyTypes
            properties = _get_properties_to_add(property_types, properties_dict)
            lat_val = lat
            long_val = long

        if last_location is None:
            locations = SearchQuery.execute(
                client, name=location_name
            ).searchForEntity.edges

            locations = [
                location.node
                for location in locations
                if location.node.entityType == "location"
                and location.node.type == location_type
                and location.node.name == location_name
            ]
            if len(locations) > 1:
                raise LocationIsNotUniqueException(
                    location_name=location_name, location_type=location_type
                )
            if len(locations) == 1:
                location_details = LocationDetailsQuery.execute(
                    client, id=locations[0].entityId
                ).location
                last_location = Location(
                    name=location_details.name,
                    id=location_details.id,
                    externalId=location_details.externalId,
                )
            else:
                add_location_input = AddLocationInput(
                    name=location_name,
                    type=client.locationTypes[location_type].id,
                    latitude=lat,
                    longitude=long,
                    properties=properties,
                    externalID=externalID,
                )

                try:
                    result = AddLocationMutation.execute(
                        client, add_location_input
                    ).__dict__[ADD_LOCATION_MUTATION_NAME]
                    client.reporter.log_successful_operation(
                        ADD_LOCATION_MUTATION_NAME, add_location_input.__dict__
                    )
                except OperationException as e:
                    raise FailedOperationException(
                        client.reporter,
                        e.err_msg,
                        e.err_id,
                        ADD_LOCATION_MUTATION_NAME,
                        add_location_input.__dict__,
                    )
                last_location = Location(
                    name=result.name, id=result.id, externalId=result.externalId
                )
        else:
            location_id = last_location.id
            result = LocationChildrenQuery.execute(client, id=location_id)
            locations = result.location.children

            locations = [
                location
                for location in locations
                if location.locationType.name == location_type
                and location.name == location_name
            ]
            if len(locations) > 1:
                raise LocationIsNotUniqueException(
                    location_name=location_name, location_type=location_type
                )
            if len(locations) == 1:
                last_location = Location(
                    name=locations[0].name,
                    id=locations[0].id,
                    externalId=locations[0].externalId,
                )
            else:
                add_location_input = AddLocationInput(
                    name=location_name,
                    type=client.locationTypes[location_type].id,
                    latitude=lat_val,
                    longitude=long_val,
                    parent=location_id,
                    properties=properties,
                    externalID=externalID,
                )
                try:
                    result = AddLocationMutation.execute(
                        client, add_location_input
                    ).__dict__[ADD_LOCATION_MUTATION_NAME]
                    client.reporter.log_successful_operation(
                        ADD_LOCATION_MUTATION_NAME, add_location_input.__dict__
                    )
                except OperationException as e:
                    raise FailedOperationException(
                        client.reporter,
                        e.err_msg,
                        e.err_id,
                        ADD_LOCATION_MUTATION_NAME,
                        add_location_input.__dict__,
                    )
                last_location = Location(
                    name=result.name, id=result.id, externalId=result.externalId
                )

    if last_location is None:
        raise LocationNotFoundException()
    return last_location


def edit_location(
    client: GraphqlClient,
    location: Location,
    new_name: str,
    new_lat: float,
    new_long: float,
    new_externalID: Optional[str],
) -> Location:

    edit_location_input = EditLocationInput(
        id=location.id,
        name=new_name,
        latitude=new_lat,
        longitude=new_long,
        properties=[],
        externalID=new_externalID,
    )

    try:
        result = EditLocationMutation.execute(client, edit_location_input).__dict__[
            EDIT_LOCATION_MUTATION_NAME
        ]
        client.reporter.log_successful_operation(
            EDIT_LOCATION_MUTATION_NAME, edit_location_input.__dict__
        )
        return Location(name=result.name, id=result.id, externalId=result.externalId)

    except OperationException as e:
        raise FailedOperationException(
            client.reporter,
            e.err_msg,
            e.err_id,
            EDIT_LOCATION_MUTATION_NAME,
            edit_location_input.__dict__,
        )
        return None


def delete_location(client: GraphqlClient, location: Location) -> None:
    deps = LocationDepsQuery.execute(client, id=location.id).location
    if len(deps.files) > 0:
        raise LocationCannotBeDeletedWithDependency(location.name, "files")
    if len(deps.children) > 0:
        raise LocationCannotBeDeletedWithDependency(location.name, "children")
    if len(deps.surveys) > 0:
        raise LocationCannotBeDeletedWithDependency(location.name, "surveys")
    if len(deps.equipments) > 0:
        raise LocationCannotBeDeletedWithDependency(location.name, "equipment")
    RemoveLocationMutation.execute(client, id=location.id)


def move_location(
    client: GraphqlClient, location_id: str, new_parent_id: Optional[str]
) -> Location:
    params = {"locationID": location_id, "parentLocationID": new_parent_id}
    try:
        result = MoveLocationMutation.execute(
            client, locationID=location_id, parentLocationID=new_parent_id
        ).__dict__[MOVE_LOCATION_MUTATION_NAME]
        client.reporter.log_successful_operation(MOVE_LOCATION_MUTATION_NAME, params)
        return Location(name=result.name, id=result.id, externalId=result.externalId)

    except OperationException as e:
        raise FailedOperationException(
            client.reporter, e.err_msg, e.err_id, MOVE_LOCATION_MUTATION_NAME, params
        )
        return None


def get_locations_by_external_id(
    client: GraphqlClient, external_id: str
) -> List[Location]:

    locations = SearchQuery.execute(client, name=external_id).searchForEntity.edges

    locations = [
        location.node
        for location in locations
        if location.node.entityType == "location"
    ]

    res = []
    for location in locations:
        location_details = LocationDetailsQuery.execute(
            client, id=location.entityId
        ).location
        if location_details.externalId == external_id:
            res.append(
                Location(
                    name=location_details.name,
                    id=location_details.id,
                    externalId=location_details.externalId,
                )
            )

    return res


def get_location_documents(client: GraphqlClient, location: Location) -> List[Document]:
    result = LocationDocumentsQuery.execute(client, id=location.id)
    files = result.location.files
    return [
        Document(
            name=file.fileName,
            id=file.id,
            parentId=location.id,
            parentEntity=ImageEntity.LOCATION,
        )
        for file in files
    ]
