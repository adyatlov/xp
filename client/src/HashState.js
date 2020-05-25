export default class HashState {
  constructor(hash) {
    if (!hash) {
      this.state = {}
      return;
    }
    let state = hash.substring(1);
    state = decodeURIComponent(state);
    this.state = JSON.parse(state);
  }
  get(field) {
    return this.state[field]
  }
  set(field, value) {
    this.state[field] = value;
  }
  toHash() {
    let hash = JSON.stringify(this.state);
    return "?" + hash;
  }
}
