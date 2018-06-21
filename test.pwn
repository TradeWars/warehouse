#include "ssc.inc"

#define RUN_TESTS
#include <YSI\y_testing>


main() {
    //
}

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