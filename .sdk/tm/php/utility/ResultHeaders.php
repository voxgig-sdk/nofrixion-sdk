<?php
declare(strict_types=1);

// Nofrixion SDK utility: result_headers

class NofrixionResultHeaders
{
    public static function call(NofrixionContext $ctx): ?NofrixionResult
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
