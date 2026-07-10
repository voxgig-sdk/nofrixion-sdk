<?php
declare(strict_types=1);

// Nofrixion SDK exists test

require_once __DIR__ . '/../nofrixion_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = NofrixionSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
