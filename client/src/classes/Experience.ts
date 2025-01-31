export default class Experience {
  id?: string;
  start = new Date();
  end = new Date();
  place = '';
  position = '';
  division = '';
  speakerId?: string;

  constructor(i?: Experience) {
    PF.H.Classes.BuildClass(this, i);
  }
  static Create(speakerId?: string): Experience {
    const item = new Experience();
    item.id = PF.H.Id.UUID();
    item.speakerId = speakerId;
    // item.name = "Новая локация"
    return item;
  }
  setPlace(item: string): void {
    this.place = item;
  }
}
