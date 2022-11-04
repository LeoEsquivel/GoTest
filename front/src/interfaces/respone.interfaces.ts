export interface MailsResponse {
    took:      number;
    timed_out: boolean;
    _shards:   Shards;
    hits:      Hits;
}

export interface Total {
    value: number
}

export interface Shards {
    total:      number;
    successful: number;
    skipped:    number;
    failed:     number;
}

export interface Hits {
    total:     Total;
    max_score: number;
    hits:      Hit[];
}

export interface Hit {
    _index:       string;
    _type:        string;
    _id:          string;
    _score:       number;
    "@timestamp": Date;
    _source:      Source;
}

export interface Source {
    Bcc:                         string;
    Cc:                          string;
    "Content-Transfer-Encoding": string;
    "Content-Type":              string;
    Date:                        string;
    From:                        string;
    "Message-ID":                string;
    "Mime-Version":              string;
    Subject:                     string;
    To:                          string;
    "X-From":                    string;
    "X-To":                      string;
    "X-bcc":                     string;
    "X-cc":                      string;
    message:                     string;
}