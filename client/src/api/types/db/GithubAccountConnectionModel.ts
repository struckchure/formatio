import { DbUserModel } from "./UserModel";

export type DbGithubAccountConnectionModel = {
  /**
   * @type string
   */
  createdAt: string;
  /**
   * @type string
   */
  deletedAt: string;
  /**
   * @type string
   */
  githubEmail: string;
  /**
   * @type string
   */
  githubId: string;
  /**
   * @type integer
   */
  githubInstallationId: number;
  /**
   * @type string
   */
  githubUsername: string;
  /**
   * @type string
   */
  id: string;
  /**
   * @type string
   */
  updatedAt: string;
  user: DbUserModel;
  /**
   * @type string
   */
  userId: string;
};
