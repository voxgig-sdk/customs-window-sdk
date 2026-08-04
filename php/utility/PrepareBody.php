<?php
declare(strict_types=1);

// CustomsWindow SDK utility: prepare_body

class CustomsWindowPrepareBody
{
    public static function call(CustomsWindowContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
