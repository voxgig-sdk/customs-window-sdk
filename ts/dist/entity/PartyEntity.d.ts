import { CustomsWindowEntityBase } from '../CustomsWindowEntityBase';
import type { CustomsWindowSDK } from '../CustomsWindowSDK';
import type { Control } from '../types';
import type { Party, PartyLoadMatch, PartyListMatch, PartyCreateData, PartyUpdateData, PartyRemoveMatch } from '../CustomsWindowTypes';
declare class PartyEntity extends CustomsWindowEntityBase<Party> {
    constructor(client: CustomsWindowSDK, entopts: any);
    make(this: PartyEntity): PartyEntity;
    load(this: any, reqmatch?: PartyLoadMatch, ctrl?: Control): Promise<PartyEntity>;
    list(this: any, reqmatch?: PartyListMatch, ctrl?: Control): Promise<PartyEntity[]>;
    create(this: any, reqdata?: PartyCreateData, ctrl?: Control): Promise<PartyEntity>;
    update(this: any, reqdata?: PartyUpdateData, ctrl?: Control): Promise<PartyEntity>;
    remove(this: any, reqmatch?: PartyRemoveMatch, ctrl?: Control): Promise<PartyEntity>;
}
export { PartyEntity };
