import Sponsor from '@/classes/Sponsor';

const SponsorsSortsLib = (() => {
  function byName(): SortModel {
    return SortModel.Create(Sponsor, PF.H.Classes.GetPropertyName(Sponsor).name);
  }
  return { byName };
})();

export default SponsorsSortsLib;
