<?php
declare(strict_types=1);

// CustomsWindow SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class CustomsWindowMakeContext
{
    public static function call(array $ctxmap, ?CustomsWindowContext $basectx): CustomsWindowContext
    {
        return new CustomsWindowContext($ctxmap, $basectx);
    }
}
