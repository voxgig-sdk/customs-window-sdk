"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PaginatedPartyListListEntity = void 0;
const CustomsWindowEntityBase_1 = require("../CustomsWindowEntityBase");
// TODO: needs Entity superclass
class PaginatedPartyListListEntity extends CustomsWindowEntityBase_1.CustomsWindowEntityBase {
    constructor(client, entopts) {
        super(client, entopts);
        this.name = 'paginated_party_list_list';
        this.name_ = 'paginated_party_list_list';
        this.Name = 'PaginatedPartyListList';
    }
    make() {
        return new PaginatedPartyListListEntity(this._client, this.entopts());
    }
}
exports.PaginatedPartyListListEntity = PaginatedPartyListListEntity;
//# sourceMappingURL=PaginatedPartyListListEntity.js.map