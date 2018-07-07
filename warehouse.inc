// Provides helper functions for working with the API via Pawn.

#include <a_samp>
#include <errors>
#include <requests>
#include <map>


// OnWarehousePlayerCreate is called after WarehousePlayerCreate finishes
forward OnWarehousePlayerCreate(playerid, bool:success, message[], Error:error, Node:result);

// OnWarehousePlayerGet is called after WarehousePlayerGetBy(Name|ID) finishes
forward OnWarehousePlayerGet(playerid, bool:success, message[], Error:error, Node:result);

// OnWarehousePlayerUpdate is called after WarehousePlayerUpdate finishes
forward OnWarehousePlayerUpdate(playerid, bool:success, message[], Error:error, Node:result);


// ParseStatus extracts success and message from Warehouse success responses:
//
// {
//     "result": {...},
//     "success": false,
//     "message": "error message"
// }
//
stock Error:ParseStatus(Node:node, &bool:success, &Node:result, message[], len = sizeof message) {
    new ret;

    ret = JsonGetBool(node, "success", success);
    if(ret) {
        return Error(1, "failed to access key 'success' in status node");
    }

    ret = JsonGetString(node, "message", message, len);
    // if(ret) {
    //     err("failed to access key 'message' in status node");
    //     return ret;
    // }

    ret = JsonGetObject(node, "result", result);
    if(ret) {
        result = Node:-1;
    }

    return NoError();
}


// -
// Interface: /store/playerCreate
// -

static Map:playerCreatePlayerID;
forward _warehouse_onPlayerCreate(Request:id, E_HTTP_STATUS:status, Node:node);

stock Error:WarehousePlayerCreate(RequestsClient:client, playerid, Node:node) {
    new Request:id = RequestJSON(
        client,
        "/store/playerCreate",
        HTTP_METHOD_POST,
        "_warehouse_onPlayerCreate",
        node
    );
    if(!IsValidRequest(id)) {
        return Error(1, "failed to send playerCreate request");
    }

    MAP_insert_val_val(playerCreatePlayerID, _:id, playerid);

    return NoError();
}
public _warehouse_onPlayerCreate(Request:id, E_HTTP_STATUS:status, Node:node) {
    new playerid = MAP_get_val_val(playerCreatePlayerID, _:id);
    MAP_remove_val(playerCreatePlayerID, _:id);

    new bool:success;
    new Node:result;
    new message[128];
    new Error:e = ParseStatus(node, success, result, message);
    if(IsError(e)) {
        e = Error(1, "failed to parse status");
    }
    if(status != HTTP_STATUS_OK) {
        e = Error(1, "status was not OK");
    }

    if(message[0] == 0) {
        message[0] = '\1';
    }

    JsonToggleGC(result, false);
    CallLocalFunction("OnWarehousePlayerCreate", "ddsdd", playerid, success, message, _:e, _:result);
    JsonToggleGC(result, true);

    return;
}


// -
// Interface: /store/playerGet
// -

static Map:playerGetPlayerID;
forward _warehouse_onPlayerGet(Request:id, E_HTTP_STATUS:status, Node:node);

stock Error:WarehousePlayerGetByName(RequestsClient:client, playerid, name[]) {
    new path[128];
    format(path, sizeof path, "/store/playerGet?name=%s", name);

    new Request:id = RequestJSON(
        client,
        path,
        HTTP_METHOD_GET,
        "_warehouse_onPlayerGet"
    );
    if(!IsValidRequest(id)) {
        return Error(1, "failed to send playerGet request");
    }

    MAP_insert_val_val(playerGetPlayerID, _:id, playerid);

    return NoError();
}

public _warehouse_onPlayerGet(Request:id, E_HTTP_STATUS:status, Node:node) {
    new playerid = MAP_get_val_val(playerGetPlayerID, _:id);
    MAP_remove_val(playerGetPlayerID, _:id);

    new bool:success;
    new Node:result;
    new message[128];
    new Error:e = ParseStatus(node, success, result, message);
    if(IsError(e)) {
        e = Error(1, "failed to parse status");
    }
    if(status != HTTP_STATUS_OK) {
        e = Error(1, "status was not OK");
    }

    if(message[0] == 0) {
        message[0] = '\1';
    }

    JsonToggleGC(result, false);
    CallLocalFunction("OnWarehousePlayerGet", "ddsdd", playerid, success, message, _:e, _:result);
    JsonToggleGC(result, true);

    return;
}


// -
// Interface: /store/playerUpdate
// -

static Map:playerUpdatePlayerID;
forward _warehouse_onPlayerUpdate(Request:id, E_HTTP_STATUS:status, Node:node);

stock Error:WarehousePlayerUpdate(RequestsClient:client, playerid, Node:node) {
    new Request:id = RequestJSON(
        client,
        "/store/playerUpdate",
        HTTP_METHOD_PATCH,
        "_warehouse_onPlayerUpdate",
        node
    );
    if(!IsValidRequest(id)) {
        return Error(1, "failed to send playerUpdate request");
    }

    // Same map strategy as with the loading data process
    MAP_insert_val_val(playerUpdatePlayerID, _:id, playerid);

    return NoError();
}

public _warehouse_onPlayerUpdate(Request:id, E_HTTP_STATUS:status, Node:node) {
    new playerid = MAP_get_val_val(playerGetPlayerID, _:id);
    MAP_remove_val(playerGetPlayerID, _:id);

    new bool:success;
    new Node:result;
    new message[128];
    new Error:e = ParseStatus(node, success, result, message);
    if(IsError(e)) {
        e = Error(1, "failed to parse status");
    }
    if(status != HTTP_STATUS_OK) {
        e = Error(1, "status was not OK");
    }

    if(message[0] == 0) {
        message[0] = '\1';
    }

    JsonToggleGC(result, false);
    CallLocalFunction("OnWarehousePlayerUpdate", "ddsdd", playerid, success, message, _:e, _:result);
    JsonToggleGC(result, true);
}
