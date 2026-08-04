<?php
declare(strict_types=1);

// CustomsWindow SDK base feature

class CustomsWindowBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(CustomsWindowContext $ctx, array $options): void {}
    public function PostConstruct(CustomsWindowContext $ctx): void {}
    public function PostConstructEntity(CustomsWindowContext $ctx): void {}
    public function SetData(CustomsWindowContext $ctx): void {}
    public function GetData(CustomsWindowContext $ctx): void {}
    public function GetMatch(CustomsWindowContext $ctx): void {}
    public function SetMatch(CustomsWindowContext $ctx): void {}
    public function PrePoint(CustomsWindowContext $ctx): void {}
    public function PreSpec(CustomsWindowContext $ctx): void {}
    public function PreRequest(CustomsWindowContext $ctx): void {}
    public function PreResponse(CustomsWindowContext $ctx): void {}
    public function PreResult(CustomsWindowContext $ctx): void {}
    public function PreDone(CustomsWindowContext $ctx): void {}
    public function PreUnexpected(CustomsWindowContext $ctx): void {}
}
