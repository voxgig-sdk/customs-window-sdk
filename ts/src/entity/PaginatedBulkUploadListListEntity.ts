
import { inspect } from 'node:util'

import { CustomsWindowEntityBase } from '../CustomsWindowEntityBase'

import type {
  CustomsWindowSDK,
} from '../CustomsWindowSDK'


import type {
  Operation,
  Context,
  Control,
} from '../types'

import type {
  PaginatedBulkUploadListList,
} from '../CustomsWindowTypes'

// TODO: needs Entity superclass
class PaginatedBulkUploadListListEntity extends CustomsWindowEntityBase<PaginatedBulkUploadListList> {

  constructor(client: CustomsWindowSDK, entopts: any) {
    super(client, entopts)
    this.name = 'paginated_bulk_upload_list_list'
    this.name_ = 'paginated_bulk_upload_list_list'
    this.Name = 'PaginatedBulkUploadListList'
  }


  make(this: PaginatedBulkUploadListListEntity) {
    return new PaginatedBulkUploadListListEntity(this._client, this.entopts())
  }







}


export {
  PaginatedBulkUploadListListEntity
}
