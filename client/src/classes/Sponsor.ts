import Banner from './Banner';

export default class Sponsor {
  id?: string;
  name = '';
  site = '';

  @PF.H.Classes.GetClassConstructor(Banner)
  banners: Banner[] = [];

  constructor(i?: Event) {
    PF.H.Classes.BuildClass(this, i);
  }

  setName(name: string) {
    this.name = name;
  }
  static GetClassName(): string {
    return 'sponsor';
  }

  static Create(): Sponsor {
    const item = new Sponsor();
    item.id = PF.H.Id.UUID();
    // item.name = "Новая локация"
    return item;
  }

  addBanner(): Banner {
    const item = Banner.Create(this.id);
    this.banners.push(item);
    return item;
  }

  removeBanner(id: string): void {
    PF.H.Classes.RemoveFromClassById(id, this.banners, []);
  }
  getBanner(): Banner {
    return Arrays.GetLast(this.banners);
  }
}
