package cronjob

import (
	"github.com/karmada-io/dashboard/pkg/dataselect"
	"github.com/karmada-io/dashboard/pkg/resource/common"
	batch "k8s.io/api/batch/v1"
)

// The code below allows to perform complex data section on []batch.CronJob

type CronJobCell batch.CronJob

func (self CronJobCell) GetProperty(name dataselect.PropertyName) dataselect.ComparableValue {
	switch name {
	case dataselect.NameProperty:
		return dataselect.StdComparableString(self.ObjectMeta.Name)
	case dataselect.CreationTimestampProperty:
		return dataselect.StdComparableTime(self.ObjectMeta.CreationTimestamp.Time)
	case dataselect.NamespaceProperty:
		return dataselect.StdComparableString(self.ObjectMeta.Namespace)
	default:
		// if name is not supported then just return a constant dummy value, sort will have no effect.
		return nil
	}
}

func ToCells(std []batch.CronJob) []dataselect.DataCell {
	cells := make([]dataselect.DataCell, len(std))
	for i := range std {
		cells[i] = CronJobCell(std[i])
	}
	return cells
}

func FromCells(cells []dataselect.DataCell) []batch.CronJob {
	std := make([]batch.CronJob, len(cells))
	for i := range std {
		std[i] = batch.CronJob(cells[i].(CronJobCell))
	}
	return std
}

func getStatus(list *batch.CronJobList) common.ResourceStatus {
	info := common.ResourceStatus{}
	if list == nil {
		return info
	}

	for _, cronJob := range list.Items {
		if cronJob.Spec.Suspend != nil && !(*cronJob.Spec.Suspend) {
			info.Running++
		} else {
			info.Failed++
		}
	}

	return info
}

func getContainerImages(cronJob *batch.CronJob) []string {
	podSpec := cronJob.Spec.JobTemplate.Spec.Template.Spec
	result := make([]string, 0)

	for _, container := range podSpec.Containers {
		result = append(result, container.Image)
	}

	return result
}
