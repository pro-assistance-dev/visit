import Session from '@/classes/Session';

class S extends PF.H.BaseStore<Session> {
  constructor() {
    super(Session, 'sessions');
  }
}

const store: S = new S();
export default store;
