/**
 * Copyright (c) 2021 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import {MigrationInterface, QueryRunner} from "typeorm";
import { columnExists } from "./helper/helper";

export class AddSoftDeletedToTeam1632908105486 implements MigrationInterface {

    public async up(queryRunner: QueryRunner): Promise<any> {
        if (!(await columnExists(queryRunner, "d_b_team", "softDeleted"))) {
            await queryRunner.query("ALTER TABLE d_b_team ADD COLUMN `softDeleted` tinyint(4) NOT NULL DEFAULT '0'");
        }
    }

    public async down(queryRunner: QueryRunner): Promise<any> {
    }

}
