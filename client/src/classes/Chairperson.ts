export default class Chairperson {
  id?: string;
  name = '';
  regalias = '';
  img = '';
  photoId?: string;
  photo: IFileInfo = new FileInfo();

  constructor(i?: Chairperson) {
    PF.H.Classes.BuildClass(this, i);
  }

  getFileInfos(): IFileInfo[] {
    return [this.photo];
  }

  removePhoto(): void {
    this.photo = new FileInfo();
    this.photoId = undefined;
  }
}
