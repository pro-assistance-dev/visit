import Place from '@/classes/Place';

class S extends PF.H.BaseStore<Place> {
  constructor() {
    super(Place, 'places');
  }
}

const store: S = new S();
export default store;
