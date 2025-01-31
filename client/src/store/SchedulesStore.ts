import Schedule from '@/classes/Schedule';

class S extends PF.H.BaseStore<Schedule> {
  constructor() {
    super(Schedule, 'schedules');
  }
}

const store: S = new S();
export default store;
