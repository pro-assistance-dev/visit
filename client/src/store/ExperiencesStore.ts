import Experience from '@/classes/Experience';

class S extends PF.H.BaseStore<Experience> {
  constructor() {
    super(Experience, 'experiences');
  }
}

const store: S = new S();
export default store;
