<?php
declare(strict_types=1);

// Nofrixion SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class NofrixionMakeContext
{
    public static function call(array $ctxmap, ?NofrixionContext $basectx): NofrixionContext
    {
        return new NofrixionContext($ctxmap, $basectx);
    }
}
