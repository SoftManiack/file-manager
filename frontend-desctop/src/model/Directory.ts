export type IDirectory = {
    uid: string,
    uidUsers: string,
    uidRoot: string,
    dateCreate: Date,
    DateUpdate: Date,
    Name: string,
    isFavorite:  boolean,
    isLink:  boolean,
    isDelete: boolean,
    path: string,
    countElement: number,
    type: string
}
