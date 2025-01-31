import User from '@/classes/User';

class S extends PF.H.BaseStore<User> {
  constructor() {
    super(User, 'users');
  }
}

const store: S = new S();
export default store;
