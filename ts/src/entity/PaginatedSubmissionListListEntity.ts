
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
  PaginatedSubmissionListList,
} from '../CustomsWindowTypes'

// TODO: needs Entity superclass
class PaginatedSubmissionListListEntity extends CustomsWindowEntityBase<PaginatedSubmissionListList> {

  constructor(client: CustomsWindowSDK, entopts: any) {
    super(client, entopts)
    this.name = 'paginated_submission_list_list'
    this.name_ = 'paginated_submission_list_list'
    this.Name = 'PaginatedSubmissionListList'
  }


  make(this: PaginatedSubmissionListListEntity) {
    return new PaginatedSubmissionListListEntity(this._client, this.entopts())
  }







}


export {
  PaginatedSubmissionListListEntity
}
