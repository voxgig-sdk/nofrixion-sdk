<?php
declare(strict_types=1);

// Nofrixion SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class NofrixionFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new NofrixionBaseFeature();
            case "test":
                return new NofrixionTestFeature();
            default:
                return new NofrixionBaseFeature();
        }
    }
}
