# Nofrixion SDK utility: make_context

from core.context import NofrixionContext


def make_context_util(ctxmap, basectx):
    return NofrixionContext(ctxmap, basectx)
