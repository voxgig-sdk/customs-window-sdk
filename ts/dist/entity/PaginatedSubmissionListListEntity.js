"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PaginatedSubmissionListListEntity = void 0;
const CustomsWindowEntityBase_1 = require("../CustomsWindowEntityBase");
// TODO: needs Entity superclass
class PaginatedSubmissionListListEntity extends CustomsWindowEntityBase_1.CustomsWindowEntityBase {
    constructor(client, entopts) {
        super(client, entopts);
        this.name = 'paginated_submission_list_list';
        this.name_ = 'paginated_submission_list_list';
        this.Name = 'PaginatedSubmissionListList';
    }
    make() {
        return new PaginatedSubmissionListListEntity(this._client, this.entopts());
    }
}
exports.PaginatedSubmissionListListEntity = PaginatedSubmissionListListEntity;
//# sourceMappingURL=PaginatedSubmissionListListEntity.js.map