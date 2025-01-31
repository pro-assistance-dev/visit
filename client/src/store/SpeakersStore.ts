import Speaker from '@/classes/Speaker';

class S extends PF.H.BaseStore<Speaker> {
  constructor() {
    super(Speaker, 'speakers');
  }
}

const store: S = new S();
export default store;
