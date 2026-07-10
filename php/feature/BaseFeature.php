<?php
declare(strict_types=1);

// Nofrixion SDK base feature

class NofrixionBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(NofrixionContext $ctx, array $options): void {}
    public function PostConstruct(NofrixionContext $ctx): void {}
    public function PostConstructEntity(NofrixionContext $ctx): void {}
    public function SetData(NofrixionContext $ctx): void {}
    public function GetData(NofrixionContext $ctx): void {}
    public function GetMatch(NofrixionContext $ctx): void {}
    public function SetMatch(NofrixionContext $ctx): void {}
    public function PrePoint(NofrixionContext $ctx): void {}
    public function PreSpec(NofrixionContext $ctx): void {}
    public function PreRequest(NofrixionContext $ctx): void {}
    public function PreResponse(NofrixionContext $ctx): void {}
    public function PreResult(NofrixionContext $ctx): void {}
    public function PreDone(NofrixionContext $ctx): void {}
    public function PreUnexpected(NofrixionContext $ctx): void {}
}
