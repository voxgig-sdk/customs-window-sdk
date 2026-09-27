"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.CustomsWindowError = void 0;
class CustomsWindowError extends Error {
    isCustomsWindowError = true;
    sdk = 'CustomsWindow';
    code;
    ctx;
    status = -1;
    // `err.notFound` rather than a magic number at every call site.
    get notFound() { return 404 === this.status; }
    constructor(code, msg, ctx) {
        super(msg);
        this.code = code;
        this.ctx = ctx;
    }
}
exports.CustomsWindowError = CustomsWindowError;
//# sourceMappingURL=CustomsWindowError.js.map