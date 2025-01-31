export default class Stream {
  id?: string;
  poster = new FileInfo();
  posterId?: string;

  constructor(i?: Stream) {
    PF.H.Classes.BuildClass(this, i);
  }
}
