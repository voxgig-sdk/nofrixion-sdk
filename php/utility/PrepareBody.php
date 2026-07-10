<?php
declare(strict_types=1);

// Nofrixion SDK utility: prepare_body

class NofrixionPrepareBody
{
    public static function call(NofrixionContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
