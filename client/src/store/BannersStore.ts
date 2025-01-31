import Banner from '@/classes/Banner';

class S extends PF.H.BaseStore<Banner> {
  constructor() {
    super(Banner, 'banners');
  }
}

const store: S = new S();
export default store;
