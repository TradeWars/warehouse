#include "warehouse.inc"

#define RUN_TESTS
#include <YSI\y_testing>


new buf[4096];

Test:ParseStatus() {
    new Node:n = JsonObject(
        "result", JsonObject(
            "key", JsonString("value")
        ),
        "success", JsonBool(true),
        "message", JsonString("success")
    );

    new bool:success;
    new Node:result;
    new message[128];

    new Error:e = ParseStatus(n, success, result, message);
    ASSERT(!IsError(e));

    ASSERT(success == true);
    ASSERT(!strcmp(message, "success"));

    new key[16];
    JsonGetString(result, "key", key);
    ASSERT(!strcmp(key, "value"));
}

new RequestsClient:testClient;

public OnScriptInit() {
    print("Doing request test");

    testClient = RequestsClient("http://localhost:7788", RequestHeaders(
        "Authorization", "cunning_fox"
    ));
    if(!IsValidRequestsClient(testClient)) {
        print("failed to create requests client");
    }

    new Error:e = WarehouseIndex(testClient);
    if(IsError(e)) {
        PrintErrors();
        Handled();
    }
}

public OnWarehouseIndex(bool:success, message[], Error:error, Node:result) {
    new Error:e = WarehousePlayerCreate(testClient, 0, JsonObject(
        "account", JsonObject(
            "name", JsonString("Southclaws"),
            "pass", JsonString("$2y$12$FY26qU4VUsT00lvv.FFGA.jMCAlHVgUatwlAuE9tf7j8rnevR0ioS"),
            "ipv4", JsonString("127.0.0.1"),
            "gpci", JsonString("0000000000000000000000000000000000000000")
        )
    ));
    if(IsError(e)) {
        PrintErrors();
        Handled();
    }
    return;
}

new PlayerObjectID[25];

public OnWarehousePlayerCreate(playerid, bool:success, message[], Error:error, Node:result) {
    if(IsError(error)) {
        PrintErrors();
        Handled();
    }

    if(!success) {
        printf("Success was false: %s", message);
        return;
    }

    JsonGetNodeString(result, PlayerObjectID);
    printf("player ObjectId: %s", PlayerObjectID);

    // Get the full record

    new Error:e = WarehousePlayerGetByName(testClient, 0, "Southclaws");
    if(IsError(e)) {
        PrintErrors();
        Handled();
    }
}

public OnWarehousePlayerGet(playerid, bool:success, message[], Error:error, Node:result) {
    if(IsError(error)) {
        PrintErrors();
        Handled();
    }

    if(!success) {
        printf("Success was false: %s", message);
        return;
    }

    JsonStringify(result, buf);
    printf("success: %s", buf);

    // Update the record

    new Error:e = WarehousePlayerUpdate(
        testClient,
        0,
        JsonObject(
            "_id", JsonString(PlayerObjectID),
            "account", JsonObject(
                "name", JsonString("Southclaws"),
                "pass", JsonString("$2y$12$FY26qU4VUsT00lvv.FFGA.jMCAlHVgUatwlAuE9tf7j8rnevR0ioS"),
                "ipv4", JsonString("192.168.1.1"),
                "gpci", JsonString("1111111111111111111111111111111111111111")
            ),
            "spawn", JsonObject(
                "posx", JsonFloat(50.0),
                "posy", JsonFloat(100.0),
                "posz", JsonFloat(3.0)
            )
        )
    );
    if(IsError(e)) {
        PrintErrors();
        Handled();
    }
}

public OnWarehousePlayerUpdate(playerid, bool:success, message[], Error:error, Node:result) {
    if(IsError(error)) {
        PrintErrors();
        Handled();
    }

    if(!success) {
        printf("Success was false: %s", message);
        return;
    }

    JsonStringify(result, buf);
    printf("success: %s", buf);
}

public OnRequestFailure(Request:id, errorCode, errorMessage[], len) {
    printf("Request %d failed: %d %s", _:id, errorCode, errorMessage);
}
