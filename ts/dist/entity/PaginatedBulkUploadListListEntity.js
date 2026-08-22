"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PaginatedBulkUploadListListEntity = void 0;
const CustomsWindowEntityBase_1 = require("../CustomsWindowEntityBase");
// TODO: needs Entity superclass
class PaginatedBulkUploadListListEntity extends CustomsWindowEntityBase_1.CustomsWindowEntityBase {
    constructor(client, entopts) {
        super(client, entopts);
        this.name = 'paginated_bulk_upload_list_list';
        this.name_ = 'paginated_bulk_upload_list_list';
        this.Name = 'PaginatedBulkUploadListList';
    }
    make() {
        return new PaginatedBulkUploadListListEntity(this._client, this.entopts());
    }
}
exports.PaginatedBulkUploadListListEntity = PaginatedBulkUploadListListEntity;
//# sourceMappingURL=PaginatedBulkUploadListListEntity.js.map