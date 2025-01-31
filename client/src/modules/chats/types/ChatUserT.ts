type Constructable = { new (...args: any[]): unknown };

export default interface ChatUserT extends Constructable {
  id?: string;
  getName: () => string;
}
