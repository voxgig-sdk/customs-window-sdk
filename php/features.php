<?php
declare(strict_types=1);

// CustomsWindow SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class CustomsWindowFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new CustomsWindowBaseFeature();
            case "test":
                return new CustomsWindowTestFeature();
            default:
                return new CustomsWindowBaseFeature();
        }
    }
}
