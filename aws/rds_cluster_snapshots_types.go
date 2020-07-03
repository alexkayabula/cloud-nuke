package aws

import (
	awsgo "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gruntwork-io/gruntwork-cli/errors"
)

type DBClusterSnapshots struct {
	Snapshots []string
}

// Name of the AWS resource
func (snapshot DBClusterSnapshots) ResourceName() string {
	return "RdsSnapshots"
}

// Names of the RDS DB Cluster Snapshots
func (snapshot DBClusterSnapshots) ResourceIdentifiers() []string {
	return snapshot.Snapshots
}

// MaxBatchSize decides how many cluster snapshots to delete in one call.
func (snapshot DBClusterSnapshots) MaxBatchSize() int {
	return 200
}

// Nuke/Delete all RDS DB Cluster snapshots
func (snapshot DBClusterSnapshots) Nuke(session *session.Session, identifiers []string) error {
	if err := nukeAllRdsClusterSnapshots(session, awsgo.StringSlice(identifiers)); err != nil {
		return errors.WithStackTrace(err)
	}

	return nil
}

type RdsClusterSnapshotDeleteError struct {
	name string
}

func (e RdsClusterSnapshotDeleteError) Error() string {
	return "RDS DB Cluster Snapshot:" + e.name + "was not deleted"
}

type RdsClusterSnapshotAvailableError struct {
	clusterName  string
	snapshotName string
}

func (e RdsClusterSnapshotAvailableError) Error() string {
	return "RDS DB Cluster Snapshot" + e.snapshotName + "not currently in available or failed state"
}

type RdsClusterAvailableError struct {
	name string
}

func (e RdsClusterAvailableError) Error() string {
	return "RDS DB Cluster " + e.name + "not in available state"
}
