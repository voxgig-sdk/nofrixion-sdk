<?php
declare(strict_types=1);

// Nofrixion SDK utility: feature_add

class NofrixionFeatureAdd
{
    public static function call(NofrixionContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
