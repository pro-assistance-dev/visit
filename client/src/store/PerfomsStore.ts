import Perfom from '@/classes/Perfom';

class S extends PF.H.BaseStore<Perfom> {
  constructor() {
    super(Perfom, 'perfoms');
  }
}

const store: S = new S();
export default store;
