package klcpermit

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_getCRDName(t *testing.T) {
	Obj := metav1.ObjectMeta{}
	t.Log(len(Obj.Annotations))
	t.Log(Obj.Annotations)
	t.Log(make(map[string]string))

	tests := []struct {
		name string
		pod  *corev1.Pod
		want string
	}{
		{
			name: "properly labeld pod",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						WorkloadAnnotation: "myworkload",
						VersionAnnotation:  "0.0.1",
						AppAnnotation:      "myapp",
					},
				},
			},
			want: "myapp-myworkload-0.0.1",
		},

		{
			name: "properly annotated pod",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						WorkloadAnnotation: "myworkload",
						VersionAnnotation:  "0.0.1",
						AppAnnotation:      "myapp",
					},

					Labels: map[string]string{
						WorkloadAnnotation: "myotherworkload",
						VersionAnnotation:  "0.0.1",
						AppAnnotation:      "mynotapp",
					}},
			},
			want: "myapp-myworkload-0.0.1",
		},

		{
			name: "annotated and labeled pod",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						WorkloadAnnotation: "myworkload",
						VersionAnnotation:  "0.0.1",
						AppAnnotation:      "myapp",
					},
				},
			},
			want: "myapp-myworkload-0.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCRDName(tt.pod); got != tt.want {
				t.Errorf("getCRDName() = %v, want %v", got, tt.want)
			}
		})
	}
}
