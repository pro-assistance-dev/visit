export default class UniqueUsersByHour {
  id?: string;
  dateHour: Date = new Date();
  constructor(i?: UniqueUsersByHour) {
    PF.H.Classes.BuildClass(this, i);
  }
}
