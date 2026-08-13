# CustomsWindow SDK utility: make_context

from projectname_sdk.core.context import CustomsWindowContext


def make_context_util(ctxmap, basectx):
    return CustomsWindowContext(ctxmap, basectx)
