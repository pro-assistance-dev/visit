import Sponsor from '@/classes/Sponsor';

class S extends PF.H.BaseStore<Sponsor> {
  constructor() {
    super(Sponsor, 'sponsors');
  }
}

const store: S = new S();
export default store;
