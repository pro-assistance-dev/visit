import UniqueUsersByHour from '@/classes/UniqueUsersByHour';
import User from '@/classes/User';
export default class UsersActivitiesTimesSum {
  id?: string;
  timeSum: Date = new Date();
  timeDay: Date = new Date();
  user = new User();
  userId?: string;
  constructor(i?: UsersActivitiesTimesSum) {
    PF.H.Classes.BuildClass(this, i);
  }
}
