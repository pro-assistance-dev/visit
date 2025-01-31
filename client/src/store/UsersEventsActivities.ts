import UserEventActivity from '@/classes/UserEventActivity';

class S extends PF.H.BaseStore<UserEventActivity> {
  constructor() {
    super(UserEventActivity, 'users');
  }
}

const store: S = new S();
export default store;
