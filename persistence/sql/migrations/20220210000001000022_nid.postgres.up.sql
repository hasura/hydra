-- Migration generated by the command below; DO NOT EDIT.
-- hydra:generate hydra migrate gen




-- hydra_oauth2_flow
ALTER TABLE "hydra_oauth2_flow" ADD COLUMN "nid" UUID;
ALTER TABLE "hydra_oauth2_flow" ADD CONSTRAINT "hydra_oauth2_flow_nid_fk_idx" FOREIGN KEY ("nid") REFERENCES "networks" ("id") ON UPDATE RESTRICT ON DELETE CASCADE;
