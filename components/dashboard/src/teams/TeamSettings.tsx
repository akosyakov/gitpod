/**
 * Copyright (c) 2021 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import { useState } from "react";
import ConfirmationModal from "../components/ConfirmationModal";
import { PageWithSubMenu } from "../components/PageWithSubMenu";
import settingsMenu from "./team-settings-menu";

function TeamSettings() {
    const [modal, setModal] = useState(false);
    const close = () => setModal(false);

    const deleteTeam = async () => {
        console.log("clicked delete");
    };

    return <div>
        <PageWithSubMenu subMenu={settingsMenu} title='General' subtitle='Manage general team settings.'>
            <h3 className="mt-12">Delete Team</h3>
            <p className="text-base text-gray-500 pb-4">Deleting this team will also remove all associated data with this team, including projects and workspaces. Deleted teams cannot be restored!</p>
            <button className="danger secondary" onClick={() => setModal(true)}>Delete Account</button>
        </PageWithSubMenu>

        <ConfirmationModal
            title="Delete Team"
            areYouSureText="You are about to permanently delete this team including all associated data with this team."
            buttonText="Delete Team"
            visible={modal}
            onClose={close}
            onConfirm={deleteTeam}
        >
            <ol className="text-gray-500 text-sm list-outside list-decimal">
                <li className="ml-5">All projects added in this team will be deleted and cannot be restored afterwards.</li>
                <li className="ml-5">All workspaces opened for projects within this team will be deleted for all team members and cannot be restored afterwards.</li>
                <li className="ml-5">All members of this team will loose access to this team, associated projects and workspaces.</li>
            </ol>
        </ConfirmationModal>
    </div>
}

export default TeamSettings;