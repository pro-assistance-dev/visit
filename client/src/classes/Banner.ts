export default class Banner {
  id?: string;
  fileInfoId?: string;
  fileInfo = new PF.C.FileInfo();
  sponsorId?: string;
  order = 0;

  constructor(i?: Banner) {
    PF.H.Classes.BuildClass(this, i);
  }

  getFileInfos() {
    return [this.fileInfo];
  }

  removeFileInfo(): void {
    this.fileInfo = new FileInfo();
    this.fileInfoId = undefined;
  }

  static Create(sponsorId?: string): Banner {
    const item = new Banner();
    item.id = PF.H.Id.UUID();
    item.sponsorId = sponsorId;
    // item.name = "Новая локация"
    return item;
  }
}
