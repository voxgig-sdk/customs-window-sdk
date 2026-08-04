
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
  PaginatedPartyListList,
} from '../CustomsWindowTypes'

// TODO: needs Entity superclass
class PaginatedPartyListListEntity extends CustomsWindowEntityBase<PaginatedPartyListList> {

  constructor(client: CustomsWindowSDK, entopts: any) {
    super(client, entopts)
    this.name = 'paginated_party_list_list'
    this.name_ = 'paginated_party_list_list'
    this.Name = 'PaginatedPartyListList'
  }


  make(this: PaginatedPartyListListEntity) {
    return new PaginatedPartyListListEntity(this._client, this.entopts())
  }







}


export {
  PaginatedPartyListListEntity
}
