<?php
declare(strict_types=1);

// Nofrixion SDK utility: result_body

class NofrixionResultBody
{
    public static function call(NofrixionContext $ctx): ?NofrixionResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
