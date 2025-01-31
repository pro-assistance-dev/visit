export default class Place {
  id?: string;
  name = '';

  constructor(i?: Event) {
    PF.H.Classes.BuildClass(this, i);
  }

  setName(name: string) {
    this.name = name;
  }
  static Create(): Place {
    const item = new Place();
    item.id = PF.H.Id.UUID();
    // item.name = "Новая локация"
    return item;
  }
  static GetClassName(): string {
    return 'place';
  }
}
