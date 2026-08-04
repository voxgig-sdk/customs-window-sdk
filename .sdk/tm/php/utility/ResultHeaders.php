<?php
declare(strict_types=1);

// CustomsWindow SDK utility: result_headers

class CustomsWindowResultHeaders
{
    public static function call(CustomsWindowContext $ctx): ?CustomsWindowResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
