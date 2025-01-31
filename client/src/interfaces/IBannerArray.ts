export default interface IBannerArray {
  id: string;
  bannerPath: string | (() => Promise<string>);
  bannerLink: string;
}
