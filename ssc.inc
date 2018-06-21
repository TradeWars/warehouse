// Provides helper functions for working with the API via Pawn.

#include <a_samp>
#include <errors>
#include <requests>


// ParseStatus extracts success and message from SSC success responses:
//
// {
//     "result": {...},
//     "success": false,
//     "message": "error message"
// }
//
Error:ParseStatus(Node:node, &bool:success, &Node:result, message[], len = sizeof message) {
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
