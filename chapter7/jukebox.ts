class JukeboxRecord {
  constructor(
    public title: string,
    public duration: number,
  ) {
    this.title = title;
    this.duration = duration;
  }

  play(recordQueue: JukeboxRecord[]) {
    console.log(`Playing ${this.title} - ${this.duration} seconds`);
    setTimeout(() => {
      const nextRecord = recordQueue.shift();
      nextRecord?.play(recordQueue);
      // }, this.duration * 1000);
    }, this.duration * 100);
  }
}

class Jukebox {
  private playQueue: JukeboxRecord[] = [];
  private records: JukeboxRecord[];
  constructor(...records: JukeboxRecord[]) {
    this.records = records;
  }
  showRecords() {
    for (let i = 0; i < this.records.length; i++) {
      console.log(`${i + 1}. ${this.records[i].title}`);
    }
  }
  addRecord(record: JukeboxRecord) {
    this.records.push(record);
  }
  removeRecord(record: JukeboxRecord) {
    const index = this.records.indexOf(record);
    if (index > -1) {
      this.records.splice(index, 1);
    }
  }
  play() {
    const nextRecord = this.playQueue.shift();
    nextRecord?.play(this.playQueue);
  }
  queue(record: JukeboxRecord) {
    this.playQueue.push(record);
  }

  userSelect(index: number) {
    const record = this.records[index - 1];
    if (record) {
      this.queue(record);
      console.log(`Added ${record.title} to queue`);
    } else {
      console.log("Invalid selection");
      this.showRecords();
    }
  }
}

let jukebox = new Jukebox(
  new JukeboxRecord("The Beatles - Yesterday", 3),
  new JukeboxRecord("The Beatles - Yellow Submarine", 2),
  new JukeboxRecord("The Beatles - Hey Jude", 4),
  new JukeboxRecord("The Beatles - Let It Be", 3),
  new JukeboxRecord("The Beatles - Come Together", 3),
);

jukebox.showRecords();
for (let i = 0; i < 50; i++) {
  jukebox.userSelect((i % 5) + 1);
}
jukebox.play();
